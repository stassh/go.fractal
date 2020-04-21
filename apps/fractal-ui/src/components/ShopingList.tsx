import * as React from "react";

export interface ShoppingListProps {
  name: string;
}

export class ShoppingList extends React.Component<ShoppingListProps, {}> {
  render() {
    return (
      <div className="shopping-list">

        <h1>Shopping List for {this.props.name}</h1>
        <ul>
          <li>Instagram</li>
          <li>WhatsApp</li>
          <li>Oculus</li>
        </ul>
      </div>
    );
  }
}
