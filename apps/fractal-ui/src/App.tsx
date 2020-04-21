import * as React from "react";
import { ShoppingList } from "./components/ShopingList";


const AsyncImageLazy = React.lazy(() => import("./components/AsyncImage"));

export interface AppProps { compiler: string; framework: string; }

export class App extends React.Component<AppProps, {}> {

    private renderLoader = () => <div className="loader"></div>;

    render() {

        const src = "https://i.imgur.com/R5TraVR.png";
        const alt = "test";


        return (
            // <main>
            //     <Switch>
            //         <Route path="/" component={About} exact />
            //         {/* <Route path="/about" component={ShoppingList} /> */}
            //     </Switch>
            // </main>

            <div>
                <ShoppingList name="New name" />
                <h1> Test app from {this.props.compiler} and {this.props.framework} !!!!!</h1>

                <React.Suspense fallback={this.renderLoader()}>
                   <AsyncImageLazy src={src} alt={alt} />
                </React.Suspense>

            </div>
        );
    };
}