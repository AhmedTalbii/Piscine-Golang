package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Load data from a CSV file
func loadData(filepath string) ([]float64, []float64, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		return nil, nil, err
	}

	var x, y []float64
	for i, row := range rows {
		if i == 0 { // Skip header
			continue
		}
		xVal, _ := strconv.ParseFloat(row[0], 64)
		yVal, _ := strconv.ParseFloat(row[1], 64)
		
		x = append(x, xVal)
		y = append(y, yVal)
		// fmt.Printf("x : %v ,y : %v,xVal : %v,yVal : %v\n",x,y,xVal,yVal)
	}
	fmt.Printf("x : %v ,y : %v\n",x,y)
	return x, y, nil
}

// Train a linear regression model using gradient descent
func trainLinearRegression(x, y []float64, lr float64, epochs int) (float64, float64) {
	m, b := 0.0, 0.0 // Initialize parameters	
	n := float64(len(x))

	for i := 0; i < epochs; i++ {
		var dm, db float64
		for j := 0; j < len(x); j++ {
			pred := m*x[j] + b
			// fmt.Printf("i: %v,j: %v,m*x[j] : %v,b : %v\n",i,j,m*x[j],b)
			error := pred - y[j]
			// fmt.Printf("i: %v,j: %v,pred : %v,y[j] : %v\n",i,j,pred,y[j])

			dm += error * x[j]
			db += error
			// fmt.Printf("i: %v,j: %v,dm : %v,db : %v\n",i,j,dm,db)

		}
		m -= lr * dm / n
		b -= lr * db / n
		// fmt.Printf("m : %v,b : %v\n",m,b)

	}
	return m, b
}

// Predict outputs for given inputs using a trained model
func predict(x []float64, m, b float64) []float64 {
	var predictions []float64
	for _, xi := range x {
		predictions = append(predictions, m*xi+b)
	}
	return predictions
}

// Calculate Mean Squared Error (MSE) between actual and predicted values
func meanSquaredError(actual, predicted []float64) float64 {
	var sum float64
	for i := range actual {
		error := actual[i] - predicted[i]
		sum += error * error
	}
	return sum / float64(len(actual))
}

// Main function
func main() {
	// Load training and test data
	trainX, trainY, _ := loadData("train_data.csv")
	testX, testY, _ := loadData("test_data.csv")

	// Train the model
	lr := 0.01 // Learning rate
	epochs := 2000
	m, b := trainLinearRegression(trainX, trainY, lr, epochs)

	// Print the trained model parameters
	fmt.Printf("Trained Model: y = %.2fx + %.2f\n", m, b)

	// Test the model
	predictions := predict(testX, m, b)
	fmt.Println("Predictions:", predictions)

	// Evaluate the model using Mean Squared Error
	mse := meanSquaredError(testY, predictions)
	fmt.Printf("Mean Squared Error: %.4f\n", mse)
}
