​# CRUD application

## Goals

The goal of this exercise is to create a command-line interface (CLI) TODO application that allows users to manage their tasks. The application should support CRUD operations (Create, Read, Update, Delete) for both users and tasks entities.

## Basic Requirements

1. User Management
    - The application should allow users to create, read, update, and delete user profiles.
    - Users should be able to provide details such as their name, email, and any other relevant information.
2. Task Management
    - The application should enable users to create, read, update, and delete tasks.
    - Tasks should have properties such as a title, description, due date, status, and any other necessary fields.
3. User-Task Relationship
    - Each task should be associated with a specific user.
    - Users should be able to view their own tasks as well as tasks assigned to other users.
4. Command-Line Interface
    - The CLI should provide a clear and intuitive syntax for entering commands and interacting with the application.
    - It should display instructions, usage guidelines, and error messages when necessary.
5. Data Storage
    - The application should persist user and task data to a file.
    - Users and tasks should be stored separately and linked appropriately.
6. Error Handling
    - The application should handle errors gracefully and provide informative error messages to guide the user.
    - It should validate user inputs, ensuring they meet the required format and constraints.
7. User-Friendly Outputs
    - The application should display task lists and user profiles in a user-friendly and readable format.
    - Task lists should provide relevant details and clearly indicate task status and due dates.

## Sample usage

Example 1

```shell 
$ ./crud get_all_tasks 
  [{"id":1,"name":"go to work","done":"true"},{"id":2,"name":"prepare dinner","done":"false"}] 
``` 

Example 2

```shell 
$ ./crud get_all_users 
  [{"id":1,"username":"user1","email":"user1@example.com"},{"id":2,"username":"user2","email":"user2@example.com"}] 
``` 

Example 3

```shell 
$ ./crud create_task task.json 
  task created with id: 3 
``` 

Example 4

```shell 
$ ./crud delete_user 2 
  user deleted 
``` 

Example 5

```shell 
$ ./crud get_user_by_id 2 
  error: user not found 
``` 
  