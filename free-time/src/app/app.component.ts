import { Component } from '@angular/core';
import { CalendarOptions } from '@fullcalendar/core'; // useful for typechecking
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent {
  title = 'FreeTime calendar';
  events: any = [
    {title: 'Meeting', date: '2023-02-02', color: '#00FF00'},
    {title: 'Dinner', date: '2023-02-11', color: '#62F1F5'},
    {title: 'Valentine\'s Day', date: '2023-02-14', color: '#FF0000'},
    {title: 'Project Due', date: '2023-02-28', color: '#675839'},
    {title: 'Exam', date: '2023-02-20', color: '#8F8F8F'},
    {title: 'Birthday', date: '2023-02-25', color: '#0000FF'}
  ]
  calendarOptions: CalendarOptions = {
    plugins: [
      interactionPlugin,
      dayGridPlugin
    ],
    initialView: 'dayGridMonth',
    editable: true,
    selectable: true,
    events: this.events
  };
}