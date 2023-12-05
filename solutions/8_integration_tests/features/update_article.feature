Feature: Update article

# TODO: uncomment after DB cleanup implementation
#  Scenario: Update one field
#    Given fixtures in "articles" table:
#      | id                                   | title      | body      |
#      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | test title | test body |
#    When I run "update_article" command with flags:
#      | id                                   | title         |
#      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | updated title |
#    Then I see in "articles" table:
#      | id                                   | title         | body      |
#      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | updated title | test body |

  Scenario: Update few fields
# TODO: implement step
#    Given fixtures in "articles" table:
#      | id                                   | title      | body      |
#      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | test title | test body |
    When I run "update_article" command with flags:
      | id                                   | title         | body         |
      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | updated title | updated body |
    Then I see in "articles" table:
      | id                                   | title         | body         |
      | dae3f7bc-ee40-4271-a38f-f70e42750bb1 | updated title | updated body |
