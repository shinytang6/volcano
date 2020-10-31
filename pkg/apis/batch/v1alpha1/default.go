package v1alpha1

const (
	DefaultReplicas = 1
	DefaultMaxRetry = 3
)

// SetDefaults_RJob sets any unspecified values to defaults.
func SetDefaults_Job(job *Job) {
	var totalReplicas int32
	for _, task := range job.Spec.Tasks {
		if task.Replicas == nil {
			task.Replicas = new(int32)
			*task.Replicas = DefaultReplicas
		}

		totalReplicas += *task.Replicas
	}

	if job.Spec.MinAvailable == nil {
		job.Spec.MinAvailable = new(int32)
		*job.Spec.MinAvailable = totalReplicas
	}

	if job.Spec.MaxRetry == nil {
		job.Spec.MaxRetry = new(int32)
		*job.Spec.MaxRetry = DefaultMaxRetry
	}
}
