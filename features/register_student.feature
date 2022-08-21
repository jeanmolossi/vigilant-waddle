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
