# README
### To run locally
npm i

Mac: `DEBUG=jobsite-api:* npm start`

Windows: `set DEBUG=jobsite-api:* & npm start`

### To See endpoint with mock data
- in browser type "localhost:8081/jobs" after running it locally
it should look something like this:
```
{
  "data": [
    {
      "type": "jobs",
      "id": "1",
      "attributes": {
        "apply_link": "apply.com/job",
        "bio": "lorem ipsum...",
        "city": "Murfreesboro",
        "company_name": "BudgetBird",
        "facebook": "facebook.com/borodev",
        "job_description": "lorem ipsum...",
        "job_type": "full-time",
        "linked_in": "linkedin.com/borodev",
        "remote": true,
        "state": "TN",
        "tech_stack": "Ember & Ruby On Rails",
        "title": "Test Title",
        "twitter": "twitter.com/borodev",
        "xp": "midLevel"
      }
    }
  ]
}
```

### Slack
To test this with Slack your user needs the `identity.basic` oAuth scope and the bot needs `chat:write`.

### Docker
#### This section might be different
Run an instance of Postgres 
`docker container run --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres -d postgres`