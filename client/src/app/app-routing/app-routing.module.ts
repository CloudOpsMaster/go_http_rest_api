import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { ProductPageComponent } from '../product-page/product-page.component'
import { ProductLayoutComponent } from '../shared/layouts/product-layout/product-layout.component'
import { RouterModule, Routes } from '@angular/router';
import { GetProductComponent } from '../get-product/get-product.component';


const routes: Routes = [
  {
      path: '', component: ProductLayoutComponent, children: [
         {path: '', redirectTo: '/products', pathMatch: 'full'},
         {path: 'products', component: ProductPageComponent},
         { path: 'product', component:  GetProductComponent}
        // {path: '', redirectTo: '/product', pathMatch: 'full'}
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
