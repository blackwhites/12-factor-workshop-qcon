# Exercise 6: Posting customer interactions to Slack

## Add a lambda function to react to secret deal codes

Now for an advaned excercise... 

* Add a `/secret?code=<user-input>` endpoint to the deals service
* Hook up the button in the frontend, e.g. via *JQuery*, to present a user prompt
* Enhance the deal service to call a lambda/fission function whenever the correct secret code was entered
* Write a lambda/fission function that reports this happy day to Slack

### Hooking up the button

You might want to add an onclick handler to the appropriate button in `topbar.html` (don't do this at home...) 
and/or add a corresponding function in `public/js/client.js` and/or `public/js/front.js`.
A `$.ajax()` Ajax call via JQuery should do the trick.
Don't forget to also add another action in `api/deals/index.js`.

Now let's rebuild the frontend image, tagged with a new version:

    docker build -t 12-factor-workshop/front-end:v3 .

Adjust the `front-end-dep.yml` we created earlier to reference the new image tag and run

    kubectl apply -f front-end-dep.yml

(hint: there's a `front-end.patch` file in `show-random-deal/` that can act as a source of inspiration)

### Add a deals endpoint in the deals service.

This is essentially the same as the default endpoint, so we can do some copy&paste to get started.
It should additionally receive a `code` parameter that is then checked.

### Writing and installing the "Lambda" function

Lets write a lambda function that triggers a Slack Webhook.
You can use this webhook here: https://hooks.slack.com/services/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

To see the webhook in action, register in [12-factor-workshop.slack.com](https://12-factor-workshop.slack.com)

Once you have written the JavaScript function, create it in fission:

    # only once needed
    fission env create --name nodejs --image fission/node-env

    # create the function
    fission function create --name slackPost --env nodejs --code slackPost.js
    
    # define a route for it
    fission route create --method GET --url /slackPost --function slackPost

Try it out manually via curl. If it works, try it in the web app.

### Add a HTTP call to the deals service

The call needs to trigger the router endpoint of fission. This is automatically available as
`router.fission` within kubernetes, so the full URL to the function is: 

    http://router.fission/slackPost

Source of inspiration can be taken from the `serverless/` directory.
You might want to rebuild the image as `v5`

    docker build -t 12-factor-workshop/deals:v5 .

And change the `image` property of the `deals-dep.yaml` again before issuing

    kubectl apply -f deals-dep.yaml