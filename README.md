Go script to scrape a popular Spanish website for flat renting. It will send an email with the obtained results.

### Example

```
./main -email=mygmailaccount -pass=mypassword -areas=tetuan,chamartin -price=650 -size=40
```
### RoadMap

- [ ] Save to csv
- [ ] Save to DB
- [X] Searching across different areas
- [X] Concurrent search 
- [X] Filter with more options
- [ ] Improve email formatting

### Dependencies

github.com/PuerkitoBio/goquery
