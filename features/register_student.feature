Feature: register student
	As a student, I want make a register in the system

	Scenario: Can not register a student if data is invalid
		When I "POST" to "/student" with:
		"""
		{
			"email": "",
			"password": ""
		}
		"""
		Then the status code received should be 400
		And the response received should match json:
		"""
		{
			"error":"Bad Request",
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

	Scenario: Can not register a student if data still invalid
		When I "POST" to "/student" with:
		"""
		{
			"email": "invalid-email",
			"password": "123"
		}
		"""
		Then the status code received should be 400
		And the response received should match json:
		"""
		{
			"error":"Bad Request",
			"errors": [
				{
					"field": "email",
					"message": "email is invalid"
				},
				{
					"field": "password",
					"message": "password must be at least 6 characters"
				}
			]
		}
		"""

	Scenario: Register user should return the user created
		When I "POST" to "/student" with:
		"""
		{
			"email": "john@doe.com",
			"password": "123456"
		}
		"""
		Then the status code received should be 201

	Scenario: Can not register student with existing email
		Given there are "users" with:
			| usr_id								| usr_email		| usr_scope	| usr_password													|
			| f61139d7-1fe6-451f-b892-c96be729ce1c	| john@doe.com	| student 	| $2a$10$sPrAnxt9E5TUzfeUl7WYge4C6gpSpnyyxJWeRg82j.eFCTthT6uKC	|
		When I "POST" to "/student" with:
		"""
		{
			"email": "john@doe.com",
			"password": "123456"
		}
		"""
		Then the status code received should be 409
		And the response received should match json:
		"""
		{"error":"email already exists"}
		"""
