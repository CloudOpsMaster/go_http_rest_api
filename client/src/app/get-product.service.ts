import { Injectable } from '@angular/core';
import { HttpClient } from '@angular/common/http';
import { Product } from './product-page/interfaces'

@Injectable({
  providedIn: 'root'
})
export class GetProductService {
  apiUrl = 'http://127.0.0.1:8080/product/3';

  constructor(private http: HttpClient) { }

  getProduct() {
    return this.http.get<Product[]>(this.apiUrl)
  }
}







