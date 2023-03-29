import { Component, VERSION } from '@angular/core';
import { CalendarOptions, EventInput } from '@fullcalendar/core'; // useful for typechecking
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import {OnInit, TemplateRef } from '@angular/core';
import { BsModalRef, BsModalService } from 'ngx-bootstrap/modal';
import { HttpClient } from '@angular/common/http'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit{
  title = 'FreeTime Calendar';
  events:EventInput[] =[];
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

  modalRef?: BsModalRef;
  constructor(private modalService: BsModalService, private httpClient: HttpClient) {}

  openModal(template: TemplateRef<any>) {
     this.modalRef = this.modalService.show(template);
  }

  ngOnInit(): void {
  }

  //to be worked on
  addEvent(){
    var event_name= document.getElementById("event_title")?.title;
    var event_color= document.getElementById("event_color");
    var event_sDate= document.getElementById("event_sDate");
    let obj1:EventInput ={title: "tests", date: '2023-02-28', color: '#33EF88'};
    this.events.push(obj1);
    this.modalRef?.hide();
  }

  //this doesn't work yet, could be useful though
  loginPage(){
    window.location.href = "loginPage.component.html";

  }

    
}