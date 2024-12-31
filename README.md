<p>This project is for storing tasks localy</p>
<p>Task-cli uses json file to store its data.</p>
Here's the project url if you are interested in:
https://roadmap.sh/projects/task-tracker

<h1>Usage</h1>
To use task-cli app just use task-cli command with its command given below:

<p>add - by this command you can add a new task to app. ensure that you have given only 1 argument to this command which is task description.</p>
<p>update - this command updates description of task. make sure to give id and new description as arguments.</p>
<p>delete - by this command you can delete a task by giving its ID. App doesn't actualy delete a task just gives a value to its "deleted at" property.</p>
<p>mark-done - this command marks status of given task as done. Simply provide task id here.</p>
<p>mark-in-progress - this command marks status of given task as in-progress. Simply provide task id here.</p>
<p>list - this command lists all tasks (excluding deleted ones). by giving 1 argument which is task status it shows tasks with the given status.</p>