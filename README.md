# EVENT SCHEDULER BOT for discord
A simple Discord bot created in Golang that allows the admin of a server to create scheduled Google Calendar events. This bot will help the admin schedule an event with auto auto-generated google meet link add attendees and feed this event to attendees's google calendars.


### Functions of bot
 * Create and schedule an event.
 * Show upcoming events.
 * Add attendees to an already scheduled event.
 * Delete an event.

### Create and schedule an event
 * To create an event, input data will be in the following format.
 * The first line starts with `!schedule` 
 * The second line will contain the summary of the event.
 * The third line will contain the start date and time of the event. `YYYY-MM-DDTHH-MM-SS`, i.e `2022-02-22T09:00:00`
 * The fourth line will contain the end date and time of the event. `YYYY-MM-DDTHH-MM-SS`, i.e `2022-02-22T11:00:00`
 * The fifth line will contain the email addresses of attendees separated by spaces. i.e `xyz@gmai.com abx@gmail.com pqr@work.ac.in`
 * ![image](https://user-images.githubusercontent.com/76701875/154680826-66aff4e4-9eea-4e65-b005-32eb50a98936.png)


### Show upcoming events.
 * To see upcoming events, enter the command `!upcoming`
 * ![image](https://user-images.githubusercontent.com/76701875/154680929-4bc6ee3c-aa61-4b84-985a-2af2103b3234.png)


### Add attendees to an already scheduled event.
 * To add attendees to an already scheduled event, first you need to see the upcoming event by the `!upcoming` command.
 * After getting the list of upcoming events, You can add any number of attendees by the following command.
 * The first line starts with `!update` 
 * The second line will contain the index of the event to which you are going to add your new attendee. i.e `5`
 * The third line will contain the email addresses of attendees separated by spaces. i.e `xyz@gmai.com abx@gmail.com pqr@work.ac.in`
 * ![image](https://user-images.githubusercontent.com/76701875/154681345-a53a6af2-6395-4435-abc8-35cc6b15910f.png)


### Delete an event.
 * To Delete an event, first you need to see upcoming events by the `!upcoming` command.
 * After getting the list of upcoming events, now you can delete any event from the list.
 * The first line starts with `!delete` 
 * The second line will contain the index of the event that you want to delete. i.e `5`
 * ![image](https://user-images.githubusercontent.com/76701875/154681457-3b2dff6f-a49f-4bb8-b5ad-59322556de2d.png)


### Google calendar.
 * Screenshot of Google Calendar.
 * ![image](https://user-images.githubusercontent.com/76701875/154682613-cdb51565-d245-4aac-a95c-b7407dca183e.png)





