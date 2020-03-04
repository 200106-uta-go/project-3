package kreate

/*
## kreate profile <profile name>
1. The following variables must be saved to the profile .yaml file.
    - A profile name.
    - The foreign cluster's name (the name of the fallback cluster)
    - The foreign cluster's IP
    - The foreign cluster's Port of Entry
    - The business application's (the client's application container) image URL
    - An array of the business application's container exposable ports (non-specific)
    - An array of the business application's endpoints (/players, /info, /static. /, ect.)
2. Created profiles will be saved to etc/kreate/
*/

func CreateProfile(name string) { }
