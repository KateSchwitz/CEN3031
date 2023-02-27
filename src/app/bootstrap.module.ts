import { NgModule } from "@angular/core";
import { CommonModule } from "@angular/common";
import { BrowserAnimationsModule } from '@angular/platform-browser/animations';
import { AccordionModule } from 'ngx-bootstrap/accordion';
import { ModalModule } from 'ngx-bootstrap/modal';

@NgModule({
    declarations: [],
    imports: [
        CommonModule, BrowserAnimationsModule,
        AccordionModule.forRoot(), ModalModule.forRoot()
    ],
    exports: []
})

export class BootstrapModule {}