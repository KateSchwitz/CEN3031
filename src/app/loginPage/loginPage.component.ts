import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import { Router } from '@angular/router';

@Component({
  selector: 'app-loginPage',
  templateUrl: './loginPage.component.html',
  styleUrls: ['./loginPage.component.css']
})
export class LoginPageComponent implements OnInit{

  username: string = "";
  password: string = "";

  constructor(private http: HttpClient, private router: Router) {}

  ngOnInit(): void {}

  save():void{
      this.username = "Kate"
      this.password = "aaaaa"
      console.log('Save Clicked');
  }

  onSubmit() {
    const headers = new HttpHeaders().set('Content-Type', 'application/json');

    const body = {
      username: this.username,
      password: this.password
    };

    this.http.post('http://localhost:8080/login', body, { headers }).subscribe((response: any) => {
      if (response) {
        this.router.navigate(['/dashboard']);
      } else {
        alert('Invalid credentials');
      }
    }, error => {
      console.log(error);
    });

    console.log('Username:', this.username);
    console.log('Password:', this.password);
  }
}