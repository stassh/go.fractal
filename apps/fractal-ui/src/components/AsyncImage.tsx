/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";

export interface AsyncImageProps { src: string; alt: string }
export interface AsyncImageState { loaded: boolean; error: boolean }

// https://i.imgur.com/R5TraVR.png

export class AsyncImage extends React.Component<AsyncImageProps, AsyncImageState> {
  props: AsyncImageProps
  state: AsyncImageState

  constructor(props: AsyncImageProps) {
    super(props);
    this.state = {
      loaded: false,
      error: false
    };
  }

  render() {

    return <img
      // className={this.props.className}

      src={this.props.src}
      alt={this.props.alt} />
  }

}

export default AsyncImage