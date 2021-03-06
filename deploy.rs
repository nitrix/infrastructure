use std::io::Error;
use std::process::Command;

struct Target<'a> {
    namespace: &'a str,
    filepath: &'a str,
}

const TARGETS: &[Target] = &[
    Target {
        namespace: "default",
        filepath: "default",
    },
    Target {
        namespace: "nitrixme",
        filepath: "nitrix.me",
    },
    Target {
        namespace: "nekohubcom",
        filepath: "nekohub.com",
    },
    Target {
        namespace: "ingress-nginx",
        filepath: "ingress.yml",
    },
];

const PRUNE_WHITELIST: &[&str] = &[
    "core/v1/ConfigMap",
    "core/v1/Endpoints",
    "core/v1/PersistentVolumeClaim",
    "core/v1/Pod",
    "core/v1/ReplicationController",
    "core/v1/Service",
    "batch/v1/Job",
    "batch/v1/CronJob",
    "apps/v1/DaemonSet",
    "apps/v1/Deployment",
    "apps/v1/ReplicaSet",
    "apps/v1/StatefulSet",
];

fn main() -> Result<(), Error> {
    for target in TARGETS {
        let mut cmd = Command::new("kubectl");

        // We're applying a new configuration.
        cmd.arg("apply");

        // That configuration comes from a file or a directory (recursive),
        // targets a specific namespace and contains everything about that namespace.
        cmd.arg("--all");
        cmd.arg("--recursive");
        cmd.args(&["--namespace", target.namespace]);
        cmd.args(&["-f", target.filepath]);

        // Attempt to prune everything.
        // They will only actually get pruned if they are no longer in use.
        cmd.arg("--prune");
        for w in PRUNE_WHITELIST {
            cmd.arg(format!("--prune-whitelist={}", w));
        }

        // Run.
        cmd.status()?;
    }

    Ok(())
}
