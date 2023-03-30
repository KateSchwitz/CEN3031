import { Component, OnInit, TemplateRef } from '@angular/core';
import { BsModalRef, BsModalService } from 'ngx-bootstrap/modal';

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit{
  constructor(private modalService: BsModalService) {}
  modalRef?: BsModalRef;


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