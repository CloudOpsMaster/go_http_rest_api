import { Component, OnInit } from '@angular/core';
import { ProductService } from '../product.service';
import { Product } from './interfaces';

@Component({
  selector: 'app-product-page',
  templateUrl: './product-page.component.html',
  styleUrls: ['./product-page.component.css']
})
export class ProductPageComponent implements OnInit {
    products$: Product[];
 
    constructor(private productService: ProductService) { }

  ngOnInit(){
  this.productService.getProducts()
    .subscribe(data => this.products$ = data);
  }
}
