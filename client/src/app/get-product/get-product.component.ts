import { Component, OnInit } from '@angular/core';
import { GetProductService } from '../get-product.service';
import { Product } from '../product-page/interfaces';

@Component({
  selector: 'app-get-product',
  templateUrl: './get-product.component.html',
  styleUrls: ['./get-product.component.css']
})
export class GetProductComponent implements OnInit {
    product$: Product[]

  constructor(private getProductService:  GetProductService ) { }

  ngOnInit() {
    this.getProductService.getProduct()
    .subscribe(data => this.product$ = data)
  }

}

