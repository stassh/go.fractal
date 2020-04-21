/*!
 * @2018 Stanislav Shabalin
 */

import * as React from "react";
import { ShoppingList } from "./ShopingList";

const AsyncImageLazy = React.lazy(() => import("./AsyncImage"));

export interface HelloProps { compiler: string; framework: string; }

// 'HelloProps' describes the shape of props.
// State is never set so we use the '{}' type.
export class Hello extends React.Component<HelloProps, {}> {

    private renderLoader = () => <div className="loader"></div>;

    render() {

        const src = "https://i.imgur.com/R5TraVR.png";
        const alt = "test";


        return (
            <div>
                <ShoppingList name="New name" />
                <h1> Test Hello from {this.props.compiler} and {this.props.framework} !!!!!</h1>

                <React.Suspense fallback={this.renderLoader()}>
                   <AsyncImageLazy src={src} alt={alt} />
                </React.Suspense>
            </div>
        );
    };
}