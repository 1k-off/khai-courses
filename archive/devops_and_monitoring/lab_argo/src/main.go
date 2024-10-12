package main

import (
	"context"
	"encoding/json"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	fiberadapter "github.com/awslabs/aws-lambda-go-api-proxy/fiber"
	"github.com/gofiber/fiber/v2"
	"io"
	"lab-argo-app/pkg/entities"
	"lab-argo-app/pkg/schedule"
	"lab-argo-app/web/routes"
	"log"
	"os"
)

var fiberLambda *fiberadapter.FiberLambda

func main() {
	dbFilePath := "./mock/schedule.json"
	dbFile, err := os.Open(dbFilePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer dbFile.Close()

	data, err := io.ReadAll(dbFile)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}
	var db entities.Schedule
	err = json.Unmarshal(data, &db)
	if err != nil {
		log.Fatalf("failed to unmarshal data: %s", err)
	}

	scheduleRepo := schedule.NewRepository(db)
	scheduleService := schedule.NewService(scheduleRepo)

	app := fiber.New()
	routes.ScheduleRouter(app, scheduleService)
	if isLambda() {
		fiberLambda = fiberadapter.New(app)
		lambda.Start(Handler)
	} else {
		log.Fatal(app.Listen(":3000"))
	}
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return fiberLambda.ProxyWithContext(ctx, req)
}

func isLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}
