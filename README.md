# Family Tree Management CLI Tool

The **Family Tree** management tool is a command-line interface (CLI) application built in GoLang that helps users organize and visualize relationships between family members. This tool allows users to add new persons, create relationships, connect to a family tree database, list relationships or persons, and retrieve names based on relationships.

## Usage
family-tree [command]


### Available Commands

- `add`: Add a new person or relationship to the family tree
- `completion`: Generate the autocompletion script for the specified shell
- `connect`: Connect to the family tree database
- `count`: Count the number of people in the family tree
- `help`: Help about any command
- `list`: List all relationships or persons
- `name`: Retrieve the name of a person based on a specified relationship


### Additional Feature: Automatic Reciprocal Relationship Creation

With this new feature, when connecting relationships in the family tree using the `connect` command, the tool will automatically create reciprocal relationships. For example:

```shell
family-tree connect --name="A" --relationship="father" --of="B"
```

This command will not only establish "A" as the father of "B" but will also add "B" as the son of "A" in the family tree.



## Database Schema Design

### Tables:

#### Person Table
| Column Name   | Data Type          | Constraints      |
| ------------- | ------------------ | ---------------- |
| id            | INTEGER            | PRIMARY KEY      |
| name          | TEXT               | NOT NULL         |
| gender        | TEXT               | DEFAULT "M"      |
| family_root   | BOOLEAN            | DEFAULT FALSE    |

#### Relationship Table
| Column Name   | Data Type          | Constraints      |
| ------------- | ------------------ | ---------------- |
| id            | INTEGER            | PRIMARY KEY      |
| type          | TEXT               | NOT NULL         |

#### Family Tree Table
| Column Name    | Data Type          | Constraints                        |
| -------------- | ------------------ | ---------------------------------- |
| id             | INTEGER            | PRIMARY KEY                        |
| person1_id     | INTEGER            | NOT NULL, FOREIGN KEY (person1_id) |
| person2_id     | INTEGER            | NOT NULL, FOREIGN KEY (person2_id) |
| relationship_id| INTEGER            | NOT NULL, FOREIGN KEY (relationship_id) |

### Relationships between Tables:
- The `person1_id` and `person2_id` columns in the `family_tree` table are foreign keys referencing the `id` column in the `person` table. This indicates the two persons involved in a relationship.
- The `relationship_id` column in the `family_tree` table is a foreign key referencing the `id` column in the `relationship` table. This defines the type of relationship between the two persons.
