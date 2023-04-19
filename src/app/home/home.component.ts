import { Component, OnInit, TemplateRef } from '@angular/core';
import { BsModalRef, BsModalService } from 'ngx-bootstrap/modal';
import { CalendarOptions, EventInput } from '@fullcalendar/core'; // useful for typechecking
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';
import { Title } from '@angular/platform-browser';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  constructor(private modalService: BsModalService) {}
  modalRef?: BsModalRef;
  events:EventInput[] =[
    
  ];
  calendarOptions: CalendarOptions = {
    plugins: [
      interactionPlugin,
      dayGridPlugin
    ],
    initialView: 'dayGridMonth',
    editable: true,
    selectable: true,
    events: this.events,
    select: function(start){
      console.log(start);
    }
  };

  openModal(template: TemplateRef<any>) {
     this.modalRef = this.modalService.show(template);
  }
  addEvent(){
    //requires logic here to link back and front end
    var event_name= document.getElementById("event_title")?.title;
    var event_color= document.getElementById("event_color");
    var event_sDate= document.getElementById("event_sDate");
    let obj1:EventInput ={title: "tests", date: '2023-02-28', color: '#33EF88'};
    let obj2:EventInput ={title: "test", date: '2023-04-10', color: '#ff1111'};
    this.events.push(obj1);
      this.events.push(obj2);
      this.modalRef?.hide();
  }
  ngOnInit() {
    
  }
}