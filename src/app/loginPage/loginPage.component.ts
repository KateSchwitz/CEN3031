import { Component, EventEmitter, Input, OnInit, Output } from '@angular/core';

@Component({
  selector: 'app-loginPage',
  templateUrl: './loginPage.component.html',
  styleUrls: ['./loginPage.component.css']
})
export class LoginPageComponent {
  constructor() {
    
  }

  ngOnInit(): void {}

  @Input() count = 0
  @Output() change = new EventEmitter()
  @Output() username = new String()
  @Output() password = new String()

  increment(): void {
    this.count++
    this.change.emit(this.count)
  }

  decrement(): void {
    this.count--
    this.change.emit(this.count)
  }

  save():void{
      this.username = "Kate"
      this.password = "aaaaa"
  }

  onSubmit() {
    
  }
}