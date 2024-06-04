<?php

// Define a function to greet users
function greet($name) {
    echo "Hello, $name!";
}

// Define a function to calculate the area of a rectangle
function calculateRectangleArea($length, $width) {
    return $length * $width;
}

// Define a function to calculate the area of a circle
function calculateCircleArea($radius) {
    return pi() * pow($radius, 2);
}

// Define a function to calculate the area of a triangle
function calculateTriangleArea($base, $height) {
    return 0.5 * $base * $height;
}

// Define a function to display the menu
function displayMenu() {
    echo "1. Greet User\n";
    echo "2. Calculate Rectangle Area\n";
    echo "3. Calculate Circle Area\n";
    echo "4. Calculate Triangle Area\n";
    echo "5. Exit\n";
}

// Main program
while (true) {
    displayMenu();
    $choice = (int) readline("Enter your choice: ");
    switch ($choice) {
        case 1:
            $name = readline("Enter your name: ");
            greet($name);
            break;
        case 2:
            $length = (float) readline("Enter the length: ");
            $width = (float) readline("Enter the width: ");
            echo "The area of the rectangle is: " . calculateRectangleArea($length, $width) . "\n";
            break;
        case 3:
            $radius = (float) readline("Enter the radius: ");
            echo "The area of the circle is: " . calculateCircleArea($radius) . "\n";
            break;
        case 4:
            $base = (float) readline("Enter the base: ");
            $height = (float) readline("Enter the height: ");
            echo "The area of the triangle is: " . calculateTriangleArea($base, $height) . "\n";
            break;
        case 5:
            echo "Exiting...\n";
            exit;
        default:
            echo "Invalid choice. Please try again.\n";
    }
}

?>
