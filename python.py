# Define a function to calculate the area of a rectangle
def calculate_rectangle_area(length, width):
    return length * width

# Define a function to calculate the area of a circle
def calculate_circle_area(radius):
    return 3.14 * (radius ** 2)

# Main program
print("Welcome to the Area Calculator!")

# Ask the user for the type of shape they want to calculate the area for
shape_type = input("Enter 'rectangle' or 'circle': ")

# Validate the user's input
while shape_type.lower() not in ['rectangle', 'circle']:
    shape_type = input("Invalid input. Please enter 'rectangle' or 'circle': ")

# Ask the user for the dimensions of the shape
if shape_type.lower() == 'rectangle':
    length = float(input("Enter the length of the rectangle: "))
    width = float(input("Enter the width of the rectangle: "))
    area = calculate_rectangle_area(length, width)
    print(f"The area of the rectangle is {area} square units.")
elif shape_type.lower() == 'circle':
    radius = float(input("Enter the radius of the circle: "))
    area = calculate_circle_area(radius)
    print(f"The area of the circle is {area} square units.")

print("Thank you for using the Area Calculator!")
