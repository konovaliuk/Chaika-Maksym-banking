Feature: Customer persistence
  In order to manipulate data using domain model
  I want to be able to get it from persistence sources

  Scenario: Create customer
    When I execute "Create" method in "CustomerRepository" with body:
      """json
      {
        "id": "eb4969e6-8de6-4b3b-b693-54776017e83d",
        "email": "example@example.com",
        "passwordHash": "$2a$12$Q6k9TTpvYB1OJA087UA6X.McZbfn/5IX4aaPaN8EC8WxaZj6YNtjK",
        "createdAt": "2022-11-30T16:00:00Z",
        "updatedAt": "2022-11-30T16:00:00Z"
      }
      """
    Then I see next records in "customers" table:
      | id                                   | email               | password_hash                                                | created_at           | updated_at           |
      | eb4969e6-8de6-4b3b-b693-54776017e83d | example@example.com | $2a$12$Q6k9TTpvYB1OJA087UA6X.McZbfn/5IX4aaPaN8EC8WxaZj6YNtjK | 2022-11-30T16:00:00Z | 2022-11-30T16:00:00Z |

  Scenario: Get customer by email
    Given the next records exist in "customers" table:
      | id                                   | email               | password_hash                                                | created_at           | updated_at           |
      | eb4969e6-8de6-4b3b-b693-54776017e83d | example@example.com | $2a$12$Q6k9TTpvYB1OJA087UA6X.McZbfn/5IX4aaPaN8EC8WxaZj6YNtjK | 2022-11-30T16:00:00Z | 2022-11-30T16:00:00Z |
    When I execute "GetByEmail" method in "CustomerProjectionDAO" with body:
      """json
      {
        "email": "example@example.com"
      }
      """
    Then customer method should return:
      """json
      {
        "id": "eb4969e6-8de6-4b3b-b693-54776017e83d",
        "email": "example@example.com",
        "passwordHash": "$2a$12$Q6k9TTpvYB1OJA087UA6X.McZbfn/5IX4aaPaN8EC8WxaZj6YNtjK",
        "createdAt": "2022-11-30T16:00:00Z",
        "updatedAt": "2022-11-30T16:00:00Z"
      }
      """

