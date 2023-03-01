import { NgModule } from '@angular/core';
import { BrowserModule } from '@angular/platform-browser';
import { FullCalendarModule } from '@fullcalendar/angular';
import dayGridPlugin from '@fullcalendar/daygrid';
import interactionPlugin from '@fullcalendar/interaction';

import { AppRoutingModule } from './app-routing.module';
import { AppComponent } from './app.component';
import { CalendarComponent } from './calendar/calendar.component';

import { CommonModule } from '@angular/common';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BootstrapModule } from './bootstrap.module';
import { AccordionModule } from 'ngx-bootstrap/accordion';
import { AlertModule,AlertConfig } from 'ngx-bootstrap/alert';
import { ButtonsModule } from 'ngx-bootstrap/buttons';
import { FormsModule } from '@angular/forms';
import { CarouselModule } from 'ngx-bootstrap/carousel';
import { CollapseModule } from 'ngx-bootstrap/collapse';
import { BsDatepickerModule, BsDatepickerConfig } from 'ngx-bootstrap/datepicker';
import { BsDropdownModule,BsDropdownConfig } from 'ngx-bootstrap/dropdown';
import { ModalModule, BsModalService } from 'ngx-bootstrap/modal';
import { LoginPageComponent } from './loginPage/loginPage.component';
import { RouterModule } from '@angular/router'; //THIS IS IMPORTANT FOR SWITCHING PAGES


@NgModule({
  declarations: [
    AppComponent,
    CalendarComponent,
    LoginPageComponent //FOR EACH PAGE CREATED, NEED TO DECLARE COMPONENTS
  ],
  imports: [
    BrowserModule,
    AppRoutingModule,
    FullCalendarModule,
    CommonModule, BrowserAnimationsModule,
    BootstrapModule, ModalModule.forRoot(),
    AccordionModule,
    AlertModule,
    ButtonsModule,
    FormsModule,
    CarouselModule,
    CollapseModule,
    BsDatepickerModule.forRoot(),
    BsDropdownModule,
    ModalModule, 
    RouterModule.forRoot([ //THESE PATHS SHOULD EVENTUALLY SWTICH PAGES
      {path: '/loginPage', component: LoginPageComponent},
      {path: '', redirectTo: '/loginPage', pathMatch:'full'}
    ])
  ],
  providers: [AlertConfig, BsDatepickerConfig, BsDropdownConfig,BsModalService],
  bootstrap: [AppComponent]
})
export class AppModule { }
