/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";
import { Rect } from "../model/Rect";
import AsyncImage from "./AsyncImage";


export interface FractalGridProps { screenRect: Rect};
export interface FractalGridState {};

export class FractalGrid extends React.Component<FractalGridProps, FractalGridState> {

  render() {

    const tileSize = 80

    const items = [];
    for (let i = 0; i < 10; i++) {
      const src = `http://localhost:8080/tile?width=${tileSize}&height=${tileSize}`;


      const alt = `test ${i}`;

      items.push(
        <div>
          <AsyncImage src={src} alt={alt}/>
        </div>
      );
    }

    const rows = [];
    for (let i = 0; i < 10; i++) {
      rows.push(
        <div className="row">
          <div className="col-10 mx-auto">
            {items}
          </div>
        </div>
      );
    }

    return(
      <div>
        <h1>FractalGrid</h1>
        <div className="container">
          {rows}
        </div>
      </div>
    );
  }
}
