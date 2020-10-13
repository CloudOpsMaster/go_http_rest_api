import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

                        
import { AppComponent } from './app.component';
import { ProductPageComponent } from './product-page/product-page.component';
import { ProductLayoutComponent } from './shared/layouts/product-layout/product-layout.component';
import { Routes, RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing/app-routing.module';

@NgModule({
  declarations: [
    AppComponent,
    ProductPageComponent,
    ProductLayoutComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    

   
   
  ],
  exports: [RouterModule],
  providers: [],
  bootstrap: [AppComponent]
})
export class AppModule { }


