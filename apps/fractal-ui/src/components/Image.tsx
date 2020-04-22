/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";

export interface ImageProps { src: string; alt: string }
export interface ImageState { loaded: boolean; error: boolean }

// https://i.imgur.com/R5TraVR.png

export class Image extends React.Component<ImageProps, ImageState> {
  constructor(props: ImageProps) {
    super(props);
    this.state = {
      loaded: false,
      error: false
    };
  }

  render() {

    return <img
      className="float-left"

      src={this.props.src}
      alt={this.props.alt}
      width="80" height="80"
       />
  }

}

export default Image