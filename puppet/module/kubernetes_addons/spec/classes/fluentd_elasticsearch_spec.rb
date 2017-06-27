require 'spec_helper'
describe 'kubernetes_addons::fluentd_elasticsearch' do
  let(:pre_condition) do
    "
      class kubernetes{}
      define kubernetes::apply(
        $manifests,
      ){}
    "
  end

  let(:manifests) do
    catalogue.resource('Kubernetes::Apply', 'fluentd-elasticsearch').send(:parameters)[:manifests]
  end

  context 'with defaults' do
    it 'be valid yaml' do
      manifests.each do |manifest|
        YAML.parse manifest
      end
    end

    it 'have image set' do
      expect(manifests[0]).to match(%{^[-\s].*image: [^:]+:[^:]+$})
    end

    it 'have resources set' do
      expect(manifests[0]).to match(%{^[-\s].*cpu: [0-9]+})
      expect(manifests[0]).to match(%{^[-\s].*memory: [0-9]+})
    end
  end
end
