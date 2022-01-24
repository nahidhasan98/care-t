from service.user import User as userService
from model.user import User

print("Program is running...")

# displaying options
print("\n==============================")
print("1. Get user list")
print("2. Create a new user")
print("==============================")

# taking user choice
option = int(input("\nPlease select an option: "))

if option == 1:  # for getting user list
    userList = userService.getUserList()

    # got data, printing now
    print("Total", len(userList), "user found. User List:")
    for user in userList:
        print(
            f'ID: {user["id"]}, Name: {user["name"]}, Email: {user["email"]}, Gender: {user["gender"]}, Status: {user["status"]}')

elif option == 2:  # creating a new user
    # taking new user data
    name = input("Enter Name: ")
    email = input("Enter Email: ")
    gender = input("Enter Gender: ")
    status = input("Enter Status (active/inactive): ")

    user = User(name, email, gender, status)
    user, err = userService.createUser(user)

    if err:  # got error
        print("Error occurred while creating user: " + err)
    else:  # no error, user created and got data
        print("\nNew user created successfully. User Details:")
        print(
            f'ID: {user["id"]}, Name: {user["name"]}, Email: {user["email"]}, Gender: {user["gender"]}, Status: {user["status"]}')
else:  # wrong option choice
    print("Wrong selection. Please try again!")

print("\nProgram exiting...")
