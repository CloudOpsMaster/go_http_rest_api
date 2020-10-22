import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';

import { HttpClientModule } from '@angular/common/http';                        
import { AppComponent } from './app.component';
import { ProductPageComponent } from './product-page/product-page.component';
import { ProductLayoutComponent } from './shared/layouts/product-layout/product-layout.component';
import { Routes, RouterModule } from '@angular/router';
import { AppRoutingModule } from './app-routing/app-routing.module';
import { ProductService } from './product.service';
import { GetProductComponent } from './get-product/get-product.component';

@NgModule({
  declarations: [
    AppComponent,
    ProductPageComponent,
    ProductLayoutComponent,
    GetProductComponent
  ],
  imports: [
    AppRoutingModule,
    BrowserModule,
    HttpClientModule
    

   
   
  ],
  exports: [RouterModule],
  providers: [ProductService],
  bootstrap: [AppComponent]
})
export class AppModule { }


