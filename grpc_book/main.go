package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "tutorial-grpc-golang/model"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &Book{})
	fmt.Println("Server listening on :50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type Book struct {
	Books []*pb.BookResponse
}

func (b *Book) CreateBook(ctx context.Context, req *pb.BookCreateRequest) (*pb.ListBooksResponse, error) {
	newBook := pb.BookResponse{
		IDBook: int64(len(b.Books) + 1),
		Title:  req.GetTitle(),
		Author: req.GetAuthor(),
	}

	b.Books = append(b.Books, &newBook)

	res := pb.ListBooksResponse{
		Messages: "Create book success",
		Code:     "201",
		Data: []*pb.BookResponse{
			&newBook,
		},
	}

	return &res, nil
}

func (b *Book) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	res := pb.ListBooksResponse{
		Messages: "List of books",
		Code:     "200",
		Data:     b.Books,
	}

	return &res, nil
}

func (b *Book) UpdateBook(ctx context.Context, req *pb.BookUpdateRequest) (*pb.ListBooksResponse, error) {
	index := getBookIndexFromID(req.IDBook,b)
	if index == -1 {
		return &pb.ListBooksResponse{
			Messages: "ID not found",
			Code:     "404",
			Data:     []*pb.BookResponse{},
		},nil
	}
	
	updateBook := pb.BookResponse{
		IDBook: req.GetIDBook(),
		Title: req.GetTitle(),
		Author: req.GetAuthor(),
	}

	b.Books[index] = &updateBook

	res := pb.ListBooksResponse{
		Messages: "Update book success",
		Code:     "200",
		Data:     []*pb.BookResponse{
			&updateBook,
		},
	}

	return &res, nil
}

func (b *Book) DeleteBook(ctx context.Context, req *pb.BookDeleteRequest) (*pb.ListBooksResponse, error) {
	index := getBookIndexFromID(req.IDBook,b)
	if index == -1 {
		return &pb.ListBooksResponse{
			Messages: "ID not found",
			Code:     "404",
			Data:     []*pb.BookResponse{},
		},nil
	}

	b.Books = append(b.Books[:index], b.Books[index+1:]...)

	return &pb.ListBooksResponse{
		Messages: "Delete book success",
		Code:     "200",
		Data:     []*pb.BookResponse{},
	},nil
}

func getBookIndexFromID(ID int64, b *Book) int {
	for i := 0; i < len(b.Books); i++ {
		if b.Books[i].IDBook == ID {
			return i
		}
	}

	return -1
}
