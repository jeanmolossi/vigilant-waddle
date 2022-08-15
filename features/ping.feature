Feature: ping application
	In order to be a healthy application
	As a resilient API
	I need to be answer it successfully

	Scenario: Ping health route and receive a pong. Game is over
		When I "GET" to "/ping"
		Then the status code received should be 200
		And the response received should match json:
		"""
		{"message":"pong"}
		"""

	Scenario: Do not accept a post in ping
		When I "POST" to "/ping"
		Then the status code received should be 405
