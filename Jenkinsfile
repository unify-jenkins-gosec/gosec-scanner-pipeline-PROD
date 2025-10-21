pipeline {
    agent any
    triggers {
        cron '00 20 * * 1-5' // Runs at 20:00 on every day-of-week from Monday through Friday
    }

    stages {
        stage('Security Scan') {
            steps {
                echo 'Running gosec security scan...'
                sh '''
                    # Install Go if not present
                    if ! command -v go >/dev/null 2>&1; then
                        GO_VERSION=1.22.3
                        curl -LO https://go.dev/dl/go$GO_VERSION.linux-amd64.tar.gz
                        mkdir -p "$HOME/go"
                        tar -C "$HOME" -xzf go$GO_VERSION.linux-amd64.tar.gz
                        export PATH="$HOME/go/bin:$PATH"
                    else
                        export PATH="$HOME/go/bin:$PATH"
                    fi

                    # Install gosec if not present
                    if ! command -v gosec >/dev/null 2>&1; then
                        go install github.com/securego/gosec/v2/cmd/gosec@latest
                        export PATH=$PATH:$(go env GOPATH)/bin
                    fi

                    # Run gosec and output SARIF
                    go mod tidy || true
                    gosec -fmt sarif -out gosec-results.sarif ./... || true
                '''
            }
        }
    }

    post {
        always {
            archiveArtifacts artifacts: 'gosec-results.sarif', fingerprint: true
        }
    }
}
