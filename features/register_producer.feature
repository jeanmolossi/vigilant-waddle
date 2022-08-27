Feature: register producer
	As a producer, I want make a register in the system

	Scenario: Can not register a producer if data is invalid
		When I "POST" to "/producer" with:
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

	Scenario: Can not register a producer if data still invalid
		When I "POST" to "/producer" with:
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
		When I "POST" to "/producer" with:
		"""
		{
			"email": "producer@doe.com",
			"password": "123456"
		}
		"""
		Then the status code received should be 201

	Scenario: Can not register producer with existing email
		Given there are "users" with:
			| usr_id								| usr_email			| usr_scope	| usr_password													|
			| f61139d7-1fe6-451f-b892-c96be729ce1d	| producer@doe.com	| producer 	| $2a$10$sPrAnxt9E5TUzfeUl7WYge4C6gpSpnyyxJWeRg82j.eFCTthT6uKC	|
		When I "POST" to "/producer" with:
		"""
		{
			"email": "producer@doe.com",
			"password": "123456"
		}
		"""
		Then the status code received should be 409
		And the response received should match json:
		"""
		{"error":"email already exists"}
		"""
