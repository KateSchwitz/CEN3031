import { Component, OnInit, TemplateRef } from '@angular/core';
import { BsModalRef, BsModalService } from 'ngx-bootstrap/modal';
import { CalendarOptions, EventInput } from '@fullcalendar/core'; // useful for typechecking
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  constructor(private modalService: BsModalService) {}
  modalRef?: BsModalRef;
  events:EventInput[] =[];
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
      this.modalRef?.hide();
  }
  ngOnInit() {
    
  }
}