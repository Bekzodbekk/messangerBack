package repository

import (
	"context"
	"message-service/genproto/messagepb"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type MessageRepo struct {
	coll *mongo.Collection
}

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Messages []Message          `bson:"messages"`
}

type Message struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	To        string             `bson:"to"`
	Message   string             `bson:"message"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
	DeletedAt *time.Time         `bson:"deleted_at,omitempty"`
}

func NewMessageRepo(mongoConn *mongo.Collection) Repository {
	return &MessageRepo{
		coll: mongoConn,
	}
}

func (r *MessageRepo) CreateMessage(ctx context.Context, req *messagepb.CreateMessageRequest) (*messagepb.CreateMessageResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", err)
	}

	newMessage := Message{
		ID:        primitive.NewObjectID(),
		To:        req.To,
		Message:   req.Message,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	filter := bson.M{"_id": userID}
	update := bson.M{"$push": bson.M{"messages": newMessage}}

	_, err = r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to create message: %v", err)
	}

	return &messagepb.CreateMessageResponse{
		Status:      true,
		MessageResp: "Message created successfully",
	}, nil
}

func (r *MessageRepo) UpdateMessage(ctx context.Context, req *messagepb.UpdateMessageRequest) (*messagepb.UpdateMessageResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", err)
	}

	messageID, err := primitive.ObjectIDFromHex(req.MessageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid message ID: %v", err)
	}

	filter := bson.M{
		"_id":          userID,
		"messages._id": messageID,
	}
	update := bson.M{
		"$set": bson.M{
			"messages.$.message":    req.Message,
			"messages.$.updated_at": time.Now(),
		},
	}

	result, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to update message: %v", err)
	}

	if result.ModifiedCount == 0 {
		return nil, status.Error(codes.NotFound, "Message not found")
	}

	return &messagepb.UpdateMessageResponse{
		Status:      true,
		MessageResp: "Message updated successfully",
	}, nil
}

func (r *MessageRepo) DeleteMessage(ctx context.Context, req *messagepb.DeleteMessageRequest) (*messagepb.DeleteMessageResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", err)
	}

	messageID, err := primitive.ObjectIDFromHex(req.MessageId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid message ID: %v", err)
	}

	filter := bson.M{
		"_id":          userID,
		"messages._id": messageID,
	}
	update := bson.M{
		"$set": bson.M{
			"messages.$.deleted_at": time.Now(),
		},
	}

	result, err := r.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Failed to delete message: %v", err)
	}

	if result.ModifiedCount == 0 {
		return nil, status.Error(codes.NotFound, "Message not found")
	}

	return &messagepb.DeleteMessageResponse{
		Status:      true,
		MessageResp: "Message deleted successfully",
	}, nil
}

func (r *MessageRepo) GetMessagesByTo(ctx context.Context, req *messagepb.GetMessagesByToRequest) (*messagepb.GetMessagesByToResponse, error) {
	userID, err := primitive.ObjectIDFromHex(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid user ID: %v", err)
	}

	filter := bson.M{
		"_id":                 userID,
		"messages.to":         req.To,
		"messages.deleted_at": bson.M{"$exists": false},
	}

	var user User
	err = r.coll.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, status.Error(codes.NotFound, "No messages found")
		}
		return nil, status.Errorf(codes.Internal, "Failed to retrieve messages: %v", err)
	}

	var messages []*messagepb.Message
	for _, msg := range user.Messages {
		if msg.To == req.To && msg.DeletedAt == nil {
			messages = append(messages, &messagepb.Message{
				Id:        msg.ID.Hex(),
				To:        msg.To,
				Message:   msg.Message,
				CreatedAt: msg.CreatedAt.Format(time.RFC3339),
				UpdatedAt: msg.UpdatedAt.Format(time.RFC3339),
			})
		}
	}

	return &messagepb.GetMessagesByToResponse{
		Status:      true,
		MessageResp: "Messages retrieved successfully",
		Messages:    messages,
	}, nil
}
