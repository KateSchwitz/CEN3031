//Angular
import { NgModule } from '@angular/core';
import { HttpClientModule } from '@angular/common/http';
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import {FormsModule} from '@angular/forms';
//Angular Routing
import { RouterModule } from '@angular/router';
import { Routes } from '@angular/router';

//Components
import { AppComponent } from './app.component';
import { HomeComponent } from './home/home.component';
import { CalendarComponent } from './calendar/calendar.component';
import { LoginPageComponent } from './loginPage/loginPage.component';
import { AboutComponent } from './about/about.component';

//FullCalendar
import { FullCalendarModule } from '@fullcalendar/angular';

//Bootstrap
import { ModalModule, BsModalService } from 'ngx-bootstrap/modal';



const routes: Routes = [
  { path: 'home', component: HomeComponent },
  { path: '', redirectTo: '/home', pathMatch: 'full'},
  {path: 'login', component:LoginPageComponent},
  {path: 'calendar', component:CalendarComponent},
  {path: 'about', component:AboutComponent}
]

@NgModule({
  declarations: [
    AppComponent,
    CalendarComponent,
    LoginPageComponent,
    HomeComponent,
    AboutComponent
    //FOR EACH PAGE CREATED, NEED TO DECLARE COMPONENTS
  ],
  imports: [
    HttpClientModule,
    FullCalendarModule,
    BrowserAnimationsModule,
    RouterModule.forRoot(routes),
  ],
  providers: [BsModalService],
  bootstrap: [AppComponent],
})

export class AppModule { }
