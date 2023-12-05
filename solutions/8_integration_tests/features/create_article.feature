Feature: Create article

  Scenario: Success
    When I run "create_article" command
    Then I see in "articles" table:
      | id                                   | title         |
      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | This is title |
