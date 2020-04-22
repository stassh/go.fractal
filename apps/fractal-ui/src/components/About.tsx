/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";
import { Rect } from "../model/Rect";

export interface AboutProps {screenRect: Rect}
export interface AboutState {}

export class About extends React.Component<AboutProps, AboutState> {
  constructor(props: AboutProps) {
    super(props)
  }
  render() {
    return (
      <div>
        <div>
          Screen size:
        </div>
        <div>
          width: {this.props.screenRect.getX2()}
        </div>
        <div>
          height: {this.props.screenRect.getY2()}
        </div>
      </div>
    );
  }
}
