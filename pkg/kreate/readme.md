# True CLI MVP

## kreate init

1. Setup Directories and Environment.
    - Makes directory for kreates mouldfolder located in /var/local/kreate
    - Makes directory for kreate’s profiles located in /etc/kreate
    - Sets up kreate's environment variable KREATE_MOULDFOLDER, value set to directory path of mouldfolder.
    - Sets up kreate’s envrionment variable KREATE_PROFILE, value set to directory path of profile.
2. Deploying Helm and Istio 
    - Istio 1.4.5 downloaded and installed.
    - Helm v2.16.3  downloaded and installed.
    - Helm initialized
    - Tiller installed into the running kubernetes cluster.
    - Istio installed with Grafana, Promethiesus, Jaeger into the running kubernetes cluster

## kreate profile <profile_name>

1. The following variables must be saved to the profile .yaml file.
    - A profile name.
    - The foreign cluster's name (the name of the fallback cluster)
    - The foreign cluster's IP
    - The foreign cluster's Port of Entry
    - The business application's (the client's application container) image URL
    - An array of the business application's container exposable ports (non-specific)
    - An array of the business application's endpoints (/players, /info, /static. /, ect.)
2. Created profiles will be saved to etc/kreate/

## kreate edit <profile_name>

1. The selected profile must be opened for modification.
2. After modification, the profile will save.

## kreate remove <profile_name> (--all, -a)

1. The specified profile must be deleted (confirmation up to implementer).
2. (up to implementer) The --all, -a flag must delete all flag.

## kreate chart <profile_name>

1. Part 1. (to be reused within kreate run)
    - The specified profile must be converted from its profile format to the helm chart's values.yaml format
    - The created values.yaml must then be copied to the kustom chart.
2. Part 2.
    - The kustom chart must then be copied to a user-friendly directory (implementation as descretioned by developer)

## kreate run <profile_name>

1. Prerequisite: Helm v2.16.3 must be installed and helm init must be ran (confirm)
2. Check: If the istio environment is already deployed to the cluster, there is no need to redeploy it.
    - If istio is not deployed, istio v1.4.5 must be deployed.
3. The custom chart must be created using Part 1. of the the 'Kreate chart' logic.
4. Check: If a custom chart is already deployed to the cluster, it must be upgraded to the new chart.
    - If a custom chart is not already deployed, the new chart must be installed rather than upgraded.
    - Possible lead for the implementing developer: <https://github.com/helm/helm/issues/3353>
