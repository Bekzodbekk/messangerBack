package repository

import (
	"context"
	"time"
	"user-service/genproto/userpb"

	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

type UserRepo struct {
	coll *mongo.Collection
	rdb  *redis.Client
}

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty"`
	FirstName string             `bson:"first_name"`
	LastName  string             `bson:"last_name"`
	Email     string             `bson:"email"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	Messages  bson.A             `bson:"messages"`
	CreatedAt string             `bson:"created_at"`
	UpdatedAt string             `bson:"updated_at"`
	DeletedAt int64              `bson:"deleted_at"`
}

func NewUserRepo(mongoConn *mongo.Collection, rdb *redis.Client) Repository {
	return &UserRepo{
		coll: mongoConn,
		rdb:  rdb,
	}
}

func (u *UserRepo) Register(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.ResponseInfo, error) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userData := map[string]interface{}{
		"first_name":      req.FirstName,
		"last_name":       req.LastName,
		"email":           req.Email,
		"username":        req.Username,
		"password":        string(passwordHash),
		"expiredPassword": GenerateExpiredPassword(),
	}

	err = u.rdb.HSet(ctx, req.Email, userData).Err()
	if err != nil {
		return nil, err
	}
	err = u.rdb.Expire(ctx, req.Email, 60*time.Second).Err()
	if err != nil {
		return nil, err
	}

	return &userpb.ResponseInfo{
		Status:  true,
		Message: "Redisga Muvaffaqiyatli Saqlandi",
	}, nil
}
func (u *UserRepo) Verify(ctx context.Context, req *userpb.VerifyRequest) (*userpb.CreateUserResponse, error) {
	result, err := u.rdb.HGet(ctx, req.Email, "expiredPassword").Result()
	if err != nil {
		return nil, err
	}

	if req.Password != result {
		return &userpb.CreateUserResponse{
			RespInfo: &userpb.ResponseInfo{
				Status:  false,
				Message: "Expired Key Xato",
			},
		}, nil
	}

	resultFullUser, err := u.rdb.HGetAll(ctx, req.Email).Result()
	if err != nil {
		return nil, err
	}

	resp, err := u.CreateUser(ctx, &userpb.CreateUserRequest{
		FirstName: resultFullUser["first_name"],
		LastName:  resultFullUser["last_name"],
		Email:     resultFullUser["email"],
		Username:  resultFullUser["username"],
		Password:  resultFullUser["password"],
	})
	if err != nil {
		return nil, err
	}

	return resp, nil
}
func (u *UserRepo) Login(ctx context.Context, req *userpb.SignInRequest) (*userpb.SignInResponse, error) {
	// Foydalanuvchini username orqali izlash
	filter := bson.M{"username": req.Username, "deleted_at": 0}
	var user User
	err := u.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userpb.SignInResponse{
				RespInfo: &userpb.ResponseInfo{
					Status:  false,
					Message: "Foydalanuvchi topilmadi",
				},
			}, nil
		}
		return nil, err
	}

	// Parolni tekshirish
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return &userpb.SignInResponse{
			RespInfo: &userpb.ResponseInfo{
				Status:  false,
				Message: "Noto'g'ri parol",
			},
		}, nil
	}

	// Muvaffaqiyatli kirish
	return &userpb.SignInResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Muvaffaqiyatli kirish",
		},
		User: &userpb.User{
			Id:        user.Id.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
			UserInfoCud: &userpb.UserInfoCUD{
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		},
	}, nil
}

func (u *UserRepo) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
	query := bson.M{
		"first_name": req.FirstName,
		"last_name":  req.LastName,
		"email":      req.Email,
		"username":   req.Username,
		"password":   req.Password,
		"messages":   bson.A{},
		"created_at": time.Now().String(),
		"updated_at": time.Now().String(),
		"deleted_at": 0,
	}
	result, err := u.coll.InsertOne(ctx, query)
	if err != nil {
		return nil, err
	}
	return &userpb.CreateUserResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Successfully created user!",
		},
		User: &userpb.User{
			Id:        result.InsertedID.(primitive.ObjectID).Hex(),
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Username:  req.Username,
			UserInfoCud: &userpb.UserInfoCUD{
				CreatedAt: query["created_at"].(string),
				UpdatedAt: query["updated_at"].(string),
				DeletedAt: int64(query["deleted_at"].(int)),
			},
		},
	}, nil
}
func (u *UserRepo) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.UpdateUserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"email":      req.Email,
			"username":   req.Username,
			"updated_at": time.Now().String(),
		},
	}

	result, err := u.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return &userpb.UpdateUserResponse{
			RespInfo: &userpb.ResponseInfo{
				Status:  false,
				Message: "User not found",
			},
		}, nil
	}

	updatedUser, err := u.GetUserById(ctx, &userpb.GetUserByIdRequest{Id: req.Id})
	if err != nil {
		return nil, err
	}

	return &userpb.UpdateUserResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Successfully updated user!",
		},
		User: updatedUser.User,
	}, nil
}
func (u *UserRepo) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.DeleteUserResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID}
	update := bson.M{
		"$set": bson.M{
			"deleted_at": time.Now().Unix(),
		},
	}

	result, err := u.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	if result.ModifiedCount == 0 {
		return &userpb.DeleteUserResponse{
			RespInfo: &userpb.ResponseInfo{
				Status:  false,
				Message: "User not found",
			},
		}, nil
	}

	return &userpb.DeleteUserResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Successfully deleted user!",
		},
	}, nil
}
func (u *UserRepo) GetUserById(ctx context.Context, req *userpb.GetUserByIdRequest) (*userpb.GetUserByIdResponse, error) {
	objID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objID, "deleted_at": 0}
	var user User
	err = u.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &userpb.GetUserByIdResponse{
				RespInfo: &userpb.ResponseInfo{
					Status:  false,
					Message: "User not found",
				},
			}, nil
		}
		return nil, err
	}

	return &userpb.GetUserByIdResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Successfully retrieved user!",
		},
		User: &userpb.User{
			Id:        user.Id.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
			UserInfoCud: &userpb.UserInfoCUD{
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		},
	}, nil
}
func (u *UserRepo) GetUserByFilter(ctx context.Context, req *userpb.GetUserByFilterRequest) (*userpb.GetUserByFilterResponse, error) {
	filter := bson.M{}
	if req.FirstName != "" {
		filter["first_name"] = req.FirstName
	}
	if req.LastName != "" {
		filter["last_name"] = req.LastName
	}
	if req.Email != "" {
		filter["email"] = req.Email
	}
	if req.Username != "" {
		filter["username"] = req.Username
	}
	filter["deleted_at"] = req.DeletedAt

	cursor, err := u.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*userpb.User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &userpb.User{
			Id:        user.Id.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
			UserInfoCud: &userpb.UserInfoCUD{
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		})
	}

	return &userpb.GetUserByFilterResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Successfully retrieved users!",
		},
		Users: users,
	}, nil
}
func (u *UserRepo) GetUsers(ctx context.Context, req *userpb.Void) (*userpb.GetUsersResponse, error) {
	filter := bson.M{"deleted_at": 0}
	cursor, err := u.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var users []*userpb.User
	for cursor.Next(ctx) {
		var user User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, &userpb.User{
			Id:        user.Id.Hex(),
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Username:  user.Username,
			UserInfoCud: &userpb.UserInfoCUD{
				CreatedAt: user.CreatedAt,
				UpdatedAt: user.UpdatedAt,
				DeletedAt: user.DeletedAt,
			},
		})
	}

	return &userpb.GetUsersResponse{
		RespInfo: &userpb.ResponseInfo{
			Status:  true,
			Message: "Successfully retrieved all users!",
		},
		User: users,
	}, nil
}
func (u *UserRepo) GetAllDirects(ctx context.Context, req *userpb.GetAllDirectsRequest) (*userpb.GetAllDirectsResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		return nil, err
	}

	var result struct {
		Messages []struct {
			ID      primitive.ObjectID `bson:"_id,omitempty"`
			To      string             `bson:"to" json:"to"`
			Message string             `bson:"message" json:"message"`
		} `bson:"messages"`
	}

	filter := bson.M{"_id": userID}
	err = u.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	var users []*userpb.User
	for _, message := range result.Messages {
		var user User
		userIdFromTo, err := primitive.ObjectIDFromHex(message.To)
		if err != nil {
			return nil, err
		}

		filter := bson.M{"_id": userIdFromTo}
		err = u.coll.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			return nil, err
		}

		// Check if the user already exists in the users slice
		exists := false
		for _, existingUser := range users {
			if existingUser.Id == user.Id.Hex() {
				exists = true
				break
			}
		}

		// If the user doesn't exist, add to the users slice
		if !exists {
			users = append(users, &userpb.User{
				Id:        user.Id.Hex(),
				FirstName: user.FirstName,
				LastName:  user.LastName,
				Email:     user.Email,
				Username:  user.Username,
				Password:  user.Password,
				UserInfoCud: &userpb.UserInfoCUD{
					CreatedAt: user.CreatedAt,
					UpdatedAt: user.UpdatedAt,
					DeletedAt: user.DeletedAt,
				},
			})
		}
	}

	return &userpb.GetAllDirectsResponse{
		Status:       true,
		DirectsCount: int64(len(users)),
		Directs:      users,
	}, nil
}

func GenerateExpiredPassword() int {
	rand.Seed(uint64(time.Now().UnixNano()))

	// 1000 dan 9999 gacha random sonni generatsiya qilish
	randomNumber := rand.Intn(9000) + 1000

	return randomNumber
}
