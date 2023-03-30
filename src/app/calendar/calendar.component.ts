import { Component, OnInit, ViewChild } from '@angular/core';
import { CalendarOptions } from '@fullcalendar/core'; // useful for typechecking
import { BsModalService, BsModalRef } from 'ngx-bootstrap/modal';
@Component({
  selector: 'app-root',
  templateUrl: './calendar.component.html',
  styleUrls: ['./calendar.component.css']
})
export class CalendarComponent implements OnInit{
  modalRef?: BsModalRef;
  title: any;
  events: any = [
    {title: 'Meeting', date: '2023-03-02', color: '#00FF00'},
    {title: 'Meeting', date: '2023-03-11', color: '#FF0000'},
    {title: 'Meeting', date: '2023-03-25', color: '#0000FF'}
  ];
  calendarOptions: CalendarOptions = {
    initialView: 'dayGridMonth',
    editable: true,
    selectable: true,
    events: this.events,
    eventClick: this.handleDateClick.bind(this)
  };
  

  config = {
    animated: true
  };
  @ViewChild('template') template!: string;
  start: any;

  constructor(private modalService: BsModalService) {

  }

  ngOnInit(): void {
    
  }

  handleDateClick(arg:any) {
    console.log(arg);
    console.log(arg.event._def.title);
    this.title = arg.event._def.title;
    this.start = arg.event.start;
    this.modalRef = this.modalService.show(this.template, this.config);
  }

  

}