import { Component, OnInit } from '@angular/core';
import { Observable } from 'rxjs';
import { map } from 'rxjs/operators';

import { Apollo } from 'apollo-angular';
import gql from 'graphql-tag';

@Component({
  selector: 'app-users',
  templateUrl: './users.component.html',
  styleUrls: ['./users.component.css']
})
export class UsersComponent implements OnInit {
  users$: Observable<any>;
  usersQuery = gql`
    query users {
      users {
      id
      firstname
      }
    }
  `;

  constructor(private apollo: Apollo) { }

  ngOnInit() {
    this.users$ = this.apollo
                      .watchQuery<any>({ query: this.usersQuery })
                      .valueChanges
                      .pipe(map(u => u.data));
  }
}
