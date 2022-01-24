import requests


class User:
    # getUserList function for getting user list
    def getUserList():
        # setting api url
        apiURL = "https://gorest.co.in/public/v1/users"

        # sending api request
        response = requests.get(apiURL)

        # taking json format response
        respData = response.json()

        # taking user list. for now ignoring 'meta' part from response
        users = respData["data"]

        return users

    # CreateUser function for creating a new user
    def createUser(user):
        # setting api url
        apiURL = "https://gorest.co.in/public/v1/users"

        # preparing payload/post data
        payload = {
            "name": user.name,
            "email": user.email,
            "gender": user.gender,
            "status": user.status
        }

        # setting up request header
        header = {
            "Authorization": "Bearer d7c01847de4c083cb154e9a533294301e9f05f93dbae7d589e42ece63226c0a3"}

        # sending api request
        response = requests.post(apiURL, data=payload, headers=header)

        # taking json format response
        respData = response.json()

        # taking user/errorList, for now ignoring 'meta' part from response
        user = respData["data"]

        # checking if error occurred or not
        if isinstance(user, list):  # returned array/list, means error occurred
            errMsg = ""
            for err in user:
                errMsg += err["field"] + " " + err["message"] + "\n"

            return (None, errMsg)

        # at this point, no error found. returning data
        return (user, None)
