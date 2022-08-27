Feature: get me data
	As a user, I want to recover my own data

	Scenario: Should receive my data complete
		Given there are "users" with:
			| usr_id								| usr_email		| usr_scope	| usr_password													|
			| f61139d7-1fe6-451f-b892-c96be729ce1c	| john1@doe.com	| student 	| $2a$10$sPrAnxt9E5TUzfeUl7WYge4C6gpSpnyyxJWeRg82j.eFCTthT6uKC	|
		Given there is user "f61139d7-1fe6-451f-b892-c96be729ce1c" logged
		When I "GET" to "/me"
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"data": {
				"id": "f61139d7-1fe6-451f-b892-c96be729ce1c",
				"email": "john1@doe.com",
				"scopes": "student"
			}
		}
		"""

	Scenario: Should receive only my ID
		Given there are "users" with:
			| usr_id								| usr_email		| usr_scope	| usr_password													|
			| f61139d7-1fe6-451f-b892-c96be729ce1c	| john1@doe.com	| student 	| $2a$10$sPrAnxt9E5TUzfeUl7WYge4C6gpSpnyyxJWeRg82j.eFCTthT6uKC	|
		Given there is user "f61139d7-1fe6-451f-b892-c96be729ce1c" logged
		When I "GET" to "/me?fields=usr_id"
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"data": {
				"id": "f61139d7-1fe6-451f-b892-c96be729ce1c"
			}
		}
		"""

	Scenario: Should receive only my Email
		Given there are "users" with:
			| usr_id								| usr_email		| usr_scope	| usr_password													|
			| f61139d7-1fe6-451f-b892-c96be729ce1c	| john1@doe.com	| student 	| $2a$10$sPrAnxt9E5TUzfeUl7WYge4C6gpSpnyyxJWeRg82j.eFCTthT6uKC	|
		Given there is user "f61139d7-1fe6-451f-b892-c96be729ce1c" logged
		When I "GET" to "/me?fields=usr_email"
		Then the status code received should be 200
		Then the response received should match json:
		"""
		{
			"data": {
				"email": "john1@doe.com"
			}
		}
		"""
