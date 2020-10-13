import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductPageComponent } from '../product-page/product-page.component'
import { ProductLayoutComponent } from '../shared/layouts/product-layout/product-layout.component'
import { RouterModule, Routes } from '@angular/router';


const routes: Routes = [
  {
      path: '', component: ProductLayoutComponent, children: [
         {path: '', redirectTo: '/products', pathMatch: 'full'},
         {path: 'products', component: ProductPageComponent},
      ]
  },
 /* {
     path: '', component: SiteLayoutComponent, canActivate:[AuthGuard], children: [
          
      ]
  } */
]



@NgModule({
  imports: [
      RouterModule.forRoot(routes)
  ],
  exports: [
     RouterModule
  ]
})

export class AppRoutingModule { }
