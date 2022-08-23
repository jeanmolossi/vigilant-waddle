Feature: login in app
	As a user, I want to do login to access the platform

	Scenario: Should receive an access token when do login
		Given there are "users" with:
			| usr_id								| usr_email		| usr_password													|
			| f61139d7-1fe6-451f-b892-c96be729ce1c	| john1@doe.com	| $2a$10$sPrAnxt9E5TUzfeUl7WYge4C6gpSpnyyxJWeRg82j.eFCTthT6uKC	|
		When I "POST" to "/auth/login" with:
		"""
		{
			"email": "john1@doe.com",
			"password": "123456"
		}
		"""
		Then the status code received should be 200
		Then the response should contain:
		"""
		{
			"access_token": "any"
		}
		"""

	Scenario: Should receive an error if try login with invalid credentials
		When I "POST" to "/auth/login" with:
		"""
		{
			"email": "john1@doe.com",
			"password": "123456789"
		}
		"""
		Then the status code received should be 401
		Then the response received should match json:
		"""
		{ "error": "invalid credentials" }
		"""

	Scenario: Should receive an error if try login without send credentials
		When I "POST" to "/auth/login" with:
		"""
		{
			"email": "",
			"password": ""
		}
		"""
		Then the status code received should be 400
		Then the response received should match json:
		"""
		{
			"error": "Bad Request",
			"errors": [
				{
					"field": "email",
					"message": "email is invalid"
				},
				{
					"field": "password",
					"message": "password is required"
				}
			]
		}
		"""

	Scenario: Should receive an error if try logout without send auth token
		When I "POST" to "/auth/logout"
		Then the status code received should be 403
		And the response received should match json:
		"""
		{"error":"forbidden"}
		"""

	Scenario: Should receive message if logout successfully
		Given there are headers:
			| Content-Type	| application/json |
			| Authorization	| @transform:toAccessToken:@db:sessions:student_id:f61139d7-1fe6-451f-b892-c96be729ce1c |
		When I "POST" to "/auth/logout"
		Then the status code received should be 202
		And the response received should match json:
		"""
		{"message":"logged out"}
		"""
