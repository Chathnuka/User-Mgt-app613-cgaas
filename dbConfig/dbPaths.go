package dbConfig

import (
	mongo "go.mongodb.org/mongo-driver/mongo"
)

var DATABASE *mongo.Database


const DATABASE_URL = "mongodb+srv://cgaas:rvyuMzkZXfLp52m7@cgaas.bbsin5h.mongodb.net/?retryWrites=true&w=majority"

const DATABASE_NAME ="User-Mgt-APP613"
