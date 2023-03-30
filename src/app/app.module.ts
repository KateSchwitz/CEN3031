//Angular
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { BrowserModule } from '@angular/platform-browser';
import {FormsModule} from '@angular/forms';
//Angular Routing
import { RouterModule } from '@angular/router';
import { Routes } from '@angular/router';

//Components
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { CalendarComponent } from './calendar/calendar.component';
import { LoginPageComponent } from './loginPage/loginPage.component';

//FullCalendar
import { FullCalendarModule } from '@fullcalendar/angular';
import { FullCalendarComponent } from '@fullcalendar/angular';

//Bootstrap
import { ModalModule, BsModalService } from 'ngx-bootstrap/modal';



const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full'},
  {path: 'login', component:LoginPageComponent},
  {path: 'calendar', component:CalendarComponent}
]

@NgModule({
  declarations: [
    AppComponent,
    CalendarComponent,
    LoginPageComponent,
    HomeComponent
    //FOR EACH PAGE CREATED, NEED TO DECLARE COMPONENTS
  ],
  imports: [
    HttpClientModule,
    BrowserModule,
    FullCalendarModule,
    FullCalendarComponent,
    BrowserAnimationsModule,
    RouterModule.forRoot(routes),
  ],
  providers: [BsModalService],
  bootstrap: [AppComponent],
})

export class AppModule { }
