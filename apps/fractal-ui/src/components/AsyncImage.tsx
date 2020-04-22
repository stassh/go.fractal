/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";

export interface AsyncImageProps { src: string; alt: string }
export interface AsyncImageState { loaded: boolean; error: boolean }

const ImageLazy = React.lazy(() => import("./Image"));

export class AsyncImage extends React.Component<AsyncImageProps, AsyncImageState> {

  private renderLoader = () => <div className="loader"><img src="https://dummyimage.com/80x80/000/fff" alt="loading"/></div>;


  constructor(props: AsyncImageProps) {
    super(props);
    this.state = {
      loaded: false,
      error: false
    };
  }

  render() {

    return (
      <React.Suspense fallback={this.renderLoader()}>
        <ImageLazy src={this.props.src} alt={this.props.alt} />
      </React.Suspense>
    )
  }
}

export default AsyncImage