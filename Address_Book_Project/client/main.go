package main

import (
	"Address_Book_Project/proto"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}

	client := proto.NewServiceClient(conn)

	g := gin.Default()
	g.GET("/add/:username/:address/:phone", func(ctx *gin.Context) {
		a := ctx.Param("username")
		b := ctx.Param("address")
		c := ctx.Param("phone")

		req := &proto.Address{Username: a, Address: b, Phone: c}
		if response, err := client.Add(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Address:", response.Address, " Username:", response.Username, " Phone:", response.Phone),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/viewall", func(ctx *gin.Context) {
		req := &proto.Empty{}

		if stream, err := client.GetAllData(ctx, req); err == nil {
			if err != nil {
				log.Fatalf("%v.GetAddress(_) = _, %v", client, err)
			}
			for {
				row, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("%v.GetAddress(_) = _, %v", client, err)
				}
				ctx.JSON(http.StatusOK, gin.H{
					"result": fmt.Sprint("Address:", row.Address, " Username:", row.Username, " Phone:", row.Phone),
				})
			}
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/viewusername/:username", func(ctx *gin.Context) {
		a := ctx.Param("username")
		req := &proto.GetUser{Username: a}

		if response, err := client.GetByUsername(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Address:", response.Address, " Username:", response.Username, " Phone:", response.Phone),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/viewaddress/:address", func(ctx *gin.Context) {
		a := ctx.Param("address")
		req := &proto.GetAddress{Address: a}

		if response, err := client.GetByAddress(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Address:", response.Address, " Username:", response.Username, " Phone:", response.Phone),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/viewphone/:phone", func(ctx *gin.Context) {
		a := ctx.Param("phone")
		req := &proto.GetPhone{Phone: a}

		if response, err := client.GetByPhone(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Address:", response.Address, " Username:", response.Username, " Phone:", response.Phone),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/update/:username/:address/:phone", func(ctx *gin.Context) {
		a := ctx.Param("username")
		b := ctx.Param("address")
		c := ctx.Param("phone")

		req := &proto.Address{Username: a, Address: b, Phone: c}
		if response, err := client.Update(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Address:", response.Address, " Username:", response.Username, " Phone:", response.Phone),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	g.GET("/delete/:username", func(ctx *gin.Context) {
		a := ctx.Param("username")

		req := &proto.GetUser{Username: a}
		if response, err := client.Delete(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprint("Address:", response.Address, " Username:", response.Username, " Phone:", response.Phone),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
